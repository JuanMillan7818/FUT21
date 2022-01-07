import React from "react";
import "../../styles/style.css";
import { PlayerList } from "../player/PlayerList";

export const ContentMain = () => {
  return (
    <>        
        <div className="main animate__animated animate__fadeIn animate__slow">
            <PlayerList />
        </div>
    </>
  );
};
