using Godot;
using System;
using LeafClient;
using Net;

public partial class BaseCharacter : CharacterBody2D
{
	public string PlayerId;
    public Vector2 SyncDir;

    private float _maxSpeed = 80.0f;
    private float _friction = 500.0f;
    private float _acceleration = 500.0f;
    private bool _isSyncStop;

    public override void _Ready()
    {
        NetManager.Instance.AddHandler(typeof(PlayerResponse), UpdateOtherPos);

        if (GlobalData.Instance.CurrentPlayerId == PlayerId)
        {
            Camera2D camera2D = new Camera2D();
            camera2D.Zoom = new Vector2(3, 3);
            AddChild(camera2D);
        }
    }

    public override void _PhysicsProcess(double delta)
    {
        if (GlobalData.Instance.CurrentPlayerId == PlayerId)
        {
            Vector2 inputDirection = new Vector2(
                Input.GetActionStrength("move_right") - Input.GetActionRawStrength("move_left"),
                Input.GetActionStrength("move_down") - Input.GetActionRawStrength("move_up")
            ).Normalized();

            Move(delta, inputDirection);
        }
        else
        {
            MoveOther(delta, SyncDir);
        }
    }

    private void Move(double delta, Vector2 dir)
    {
        Vector2 velocity = Velocity;

        if (dir != Vector2.Zero)
        {
            velocity = velocity.MoveToward(dir * (int)_maxSpeed, (float)(_acceleration * delta));

            _isSyncStop = false;
            PlayerRequest req = new PlayerRequest
            {
                PlayerId = PlayerId,
                Vel = new Velocity { X = dir.X, Y = dir.Y }
            };
            NetManager.Instance.SendMessage(req);
        }
        else
        {
            velocity = velocity.MoveToward(Vector2.Zero, (float)(_friction * delta));
            
            if (!_isSyncStop)
            {
                _isSyncStop = true;
                PlayerRequest req = new PlayerRequest
                {
                    PlayerId = PlayerId,
                    Vel = new Velocity { X = 0, Y = 0 }
                };
                NetManager.Instance.SendMessage(req);
            }
        }

        Velocity = velocity;
        MoveAndSlide();
    }
    
    private void MoveOther(double delta, Vector2 dir)
    {
        Vector2 velocity = Velocity;

        if (dir != Vector2.Zero)
        {
            velocity = velocity.MoveToward(dir * (int)_maxSpeed, (float)(_acceleration * delta));
        }
        else
        {
            velocity = velocity.MoveToward(Vector2.Zero, (float)(_friction * delta));
        }

        Velocity = velocity;
        MoveAndSlide();
    }

    private void UpdateOtherPos(object obj)
    {
        PlayerResponse resp = (PlayerResponse)obj;
        if (resp != null)
        {
            if (resp.PlayerId != PlayerId)
            {
                GlobalData.Instance.Players[resp.PlayerId].SyncDir = new Vector2(resp.Vel.X, resp.Vel.Y);
            }
        }
    }
}
