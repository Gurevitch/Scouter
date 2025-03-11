import React, { useState } from "react";
import PlayerStats from "./components/PlayerStats";

function App() {
  const [playerId, setPlayerId] = useState(null); // Track the selected playerId

  return (
    <div>
      <h1>Player Stats Dashboard</h1>

      {/* Buttons for selecting different players */}
      <div>
        <button onClick={() => setPlayerId(4)}>Player 1</button>
        <button onClick={() => setPlayerId(5)}>Player 2</button>
        <button onClick={() => setPlayerId(6)}>Player 3</button>
        {/* Add more buttons if needed */}
      </div>

      {/* Display PlayerStats Component with selected playerId */}
      {playerId && <PlayerStats playerId={playerId} />}
    </div>
  );
}

export default App;
