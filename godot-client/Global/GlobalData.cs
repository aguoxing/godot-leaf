using System;
using System.Collections.Generic;

public class GlobalData
{
    private static readonly Lazy<GlobalData> _instance = new(() => new GlobalData());

    public static GlobalData Instance => _instance.Value;

    public Dictionary<string, BaseCharacter> Players = new Dictionary<string, BaseCharacter>();

    public string CurrentPlayerId;
    public string CurrentRoomId;

    public GlobalData()
    {
    }
    
}