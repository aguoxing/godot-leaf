using Godot;
using System;
using Net;

public partial class ServerManager : Node
{
	public override void _Ready()
	{
		NetManager.Instance.Start();
	}

	public override void _Process(double delta)
	{
		NetManager.Instance.Update();
	}
}
