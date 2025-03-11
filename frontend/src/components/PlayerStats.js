import React, { useEffect, useState } from "react";

const PlayerStats = ({ playerId }) => {
    const [playerData, setPlayerData] = useState(null);

    useEffect(() => {
        // Fetch player stats when playerId is available
        if (playerId) {
            fetch(`http://localhost:8080/player/${playerId}/stats`)
                .then((response) => response.json())
                .then((data) => setPlayerData(data))
                .catch((error) => console.error("Error fetching player stats:", error));
        }
    }, [playerId]);

    // Render loading or error if no data is available
    if (!playerData) {
        return <div>Loading...</div>;
    }

    return (
        <div>
            <h2>Player Stats for {playerData.name}</h2>
            <div>
                <h3>Player Info:</h3>
                <p><strong>Age:</strong> {playerData.age}</p>
                <p><strong>Position:</strong> {playerData.position}</p>
                <p><strong>Nationality:</strong> {playerData.nationality}</p>
                <p><strong>Wages:</strong> ${playerData.wages}</p>
            </div>

            <h3>Season Stats:</h3>
            {playerData.season_stats.map((season, index) => (
                <div key={index}>
                    <h4>Season Year: {season.season_year}</h4>
                    <p><strong>Goals:</strong> {season.goals}</p>
                    <p><strong>Assists:</strong> {season.assists}</p>
                    <div>
                        <p><strong>Yellow Cards:</strong> {season.cards.yellow_card}</p>
                        <p><strong>Red Cards:</strong> {season.cards.red_card}</p>
                    </div>
                </div>
            ))}
        </div>
    );
};

export default PlayerStats;
