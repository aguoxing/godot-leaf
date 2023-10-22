using System;
using Godot;
using LeafClient;
using Net;

public partial class Main : Control
{
    private SceneManager _sceneManager;
    private RandomNumberGenerator _rng;
    private LineEdit _roomId;

    public override void _Ready()
    {
        _rng = new RandomNumberGenerator();
        _sceneManager = GetNode<SceneManager>("/root/SceneManager");
        _roomId = GetNode<LineEdit>("UI/Menus/RoomId");

        NetManager.Instance.AddHandler(typeof(PlayerResponse), NewGameCallback);
        NetManager.Instance.AddHandler(typeof(StartGame), StartGameCallback);
    }

    public override void _Process(double delta)
    {
    }

    private void OnCreateOrJoinPressed()
    {
        if (_roomId.Text != null && _roomId.Text.Trim() != "")
        {
            string playerId = Guid.NewGuid().ToString("N");

            NetManager.Instance.SendConnect();
            NetManager.Instance.SendMessage(new NewGame() { RoomId = _roomId.Text, PlayerId = playerId });

            GlobalData.Instance.CurrentPlayerId = playerId;
            GlobalData.Instance.CurrentRoomId = _roomId.Text;
        }
    }

    private void OnStartPressed()
    {
        GD.Print("===start===");
        NetManager.Instance.SendMessage(new StartGame() { RoomId = _roomId.Text });
    }

    private void OnExitPressed()
    {
        GetTree().Quit();
    }

    private void NewGameCallback(Object obj)
    {
        PlayerResponse resp = (PlayerResponse)obj;
        if (resp != null)
        {
            GD.Print(resp.Msg);
        }
    }

    private void StartGameCallback(Object obj)
    {
        StartGame resp = (StartGame)obj;
        if (resp != null)
        {
            for (int i = 0; i < resp.PlayerList.Count; i++)
            {
                PlayerInfo playerInfo = resp.PlayerList[i];
                BaseCharacter player;
                if (i % 2 == 0)
                {
                    player = GD.Load<PackedScene>("res://Player/Player1.tscn").Instantiate<BaseCharacter>();
                }
                else
                {
                    player = GD.Load<PackedScene>("res://Player/Player2.tscn").Instantiate<BaseCharacter>();
                }
                
                player.PlayerId = playerInfo.PlayerId;
                player.Position = new Vector2(playerInfo.X, playerInfo.Y);
                GlobalData.Instance.Players.Add(playerInfo.PlayerId, player);
            }

            GD.Print("goto");
            _sceneManager.GotoScene("res://Scenes/World.tscn");
        }
    }
}