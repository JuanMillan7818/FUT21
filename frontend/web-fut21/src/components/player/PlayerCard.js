import React from "react";

export const PlayerCard = ({
  firstName,
  lastName,
  positionFull,
  nation,
  club,
  img,
}) => {
  const classImg = `img-player img${img}`;
  return (
    <>
      <div className="card" data-name={lastName}>
        <div className={classImg}></div>
        <div className="content-player">
          <h2>{firstName}</h2>
          <p>{lastName}</p>
        </div>
      </div>
    </>
  );
};
