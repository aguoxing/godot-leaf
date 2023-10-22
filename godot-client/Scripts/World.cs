using System;
using Godot;
using LeafClient;
using Net;

public partial class World : Node2D
{
    private Node2D _room;

    public override void _Ready()
    {
        _room = GetNode<Node2D>("Room");
        HandlePlayerJoin();

        NetManager.Instance.AddHandler(typeof(ExitGame), HandlePlayerExit);
    }

    public override void _Process(double delta)
    {
    }

    public override void _ExitTree()
    {
        NetManager.Instance.SendMessage(new ExitGame()
            { RoomId = GlobalData.Instance.CurrentRoomId, PlayerId = GlobalData.Instance.CurrentPlayerId });
    }

    private void HandlePlayerJoin()
    {
        foreach (var keyValuePair in GlobalData.Instance.Players)
        {
            _room.AddChild(keyValuePair.Value);
        }
    }

    private void HandlePlayerExit(Object obj)
    {
        ExitGame resp = (ExitGame)obj;
        if (resp != null)
        {
            GlobalData.Instance.Players[resp.PlayerId].QueueFree();
            GlobalData.Instance.Players.Remove(resp.PlayerId);
        }
    }
}