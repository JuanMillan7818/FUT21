import React, { useContext } from "react";
import { useNavigate } from "react-router-dom";
import { ApiContext } from "../context/ApiContext";

export const PlayerCard = ({
  firstName,
  lastName,
  positionFull,
  nation,
  club,
  img,
}) => {
  const navigate = useNavigate();
  const { setSearchExpand } = useContext(ApiContext);
  const classImg = `img-player img${img}`; // manejar imagen generica de los jugadores mediante un contador

  // Renderizar el componente para expandir y ver la informacion del jugador
  const handleExpandData = () => {
    setSearchExpand({firstName, lastName, positionFull, nation, club, classImg})
    navigate(`/search`);
  };
  return (
    <>
      <div className="card animate__animated animate__fadeIn animate__fast" data-name={lastName} onClick={handleExpandData}>
        <div className={classImg}></div>
        <div className="content-player">
          <h2>{firstName}</h2>
          <p>{lastName}</p>
        </div>
      </div>
    </>
  );
};
