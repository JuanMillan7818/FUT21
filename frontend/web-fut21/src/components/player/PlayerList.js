import React, { useContext, useState } from "react";
import { ApiContext } from "../context/ApiContext";
import { PlayerCard } from "./PlayerCard";

export const PlayerList = () => {
  const {players, totalPage, url, setUrl} = useContext(ApiContext)
  const [pageActive, setPageActive] = useState(0);
  const [cantPage, setCantPage] = useState(0);

  const handleItemActive = (element, page) => {
    let lastPage = document.getElementById(pageActive);
    lastPage.className = "";

    element.target.className = "active";
    setPageActive(page);

    setUrl(url.split('=')[0]+'='+(page+1+cantPage))
  };

  const handleChangePage = (type) => {
    console.log(pageActive);
    if (type === "back") {
      if(cantPage !== 0) setCantPage(cantPage - 6);
    } else {
      if(cantPage <= totalPage ) setCantPage(cantPage + 6);
    }
  };

  let cont = 0;
  return (
    <>
      <div className="wrapper">
        {players.map((player) => {
          if (cont === 4) cont = 1;
          else cont++;
          return <PlayerCard key={player.player_id} {...player} img={cont} />;
        })}
      </div>
      <div className="pages">
        <div className="pagination">
          <ul>
            <li id="back" onClick={(e) => handleChangePage("back")}>
              &laquo;
            </li>
            {[...Array(6)].map((x, index) => {
              return (
                <li
                  key={index}
                  className={index === 0 ? "active" : ""}
                  onClick={(e) => handleItemActive(e, index)}
                  id={index}
                >
                  {index + 1 + cantPage}
                </li>
              );
            })}
            <li id="next" onClick={(e) => handleChangePage("next")}>
              &raquo;
            </li>
          </ul>
        </div>
      </div>
    </>
  );
};
