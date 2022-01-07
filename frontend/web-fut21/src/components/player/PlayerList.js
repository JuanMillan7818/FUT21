import React, { useContext, useState } from "react";
import { ApiContext } from "../context/ApiContext";
import { Warning } from "../no_find/Warning";
import { PlayerCard } from "./PlayerCard";

export const PlayerList = () => {
  const {
    players,
    totalPage,
    url,
    setUrl,
    params: { method, body },
    setParams,
  } = useContext(ApiContext);
  const [pageActive, setPageActive] = useState(0);
  const [cantPage, setCantPage] = useState(0);
  const [orderBy, setOrderBy] = useState("asc");

  const handleItemActive = (element, page) => {
    let lastPage = document.getElementById(pageActive);
    lastPage.className = "";

    element.target.className = "active";
    setPageActive(page);

    if (method === "GET") {
      let dataUrl = url.split("=");
      if (dataUrl.length > 3) {
        setUrl(
          `${dataUrl[0]}=${page + 1 + cantPage}&order=${orderBy}&search=${
            dataUrl[3]
          }`
        );
      } else {
        setUrl(`${dataUrl[0]}=${page + 1 + cantPage}&order=${orderBy}`);
      }
    } else {
      let tmp = JSON.parse(body);
      tmp = { ...tmp, Page: page + 1 + cantPage };
      setParams({
        method,
        body: JSON.stringify(tmp),
      });
    }
  };

  const handleChangePage = (type) => {
    if (type === "back") {
      if (cantPage !== 0) setCantPage(cantPage - 6);
    } else {
      if (cantPage <= totalPage) setCantPage(cantPage + 6);
    }
  };

  const handleChangeOrder = (e) => {
    let orderby = "";
    if (e.target.className === "order asc") {
      e.target.className = "order desc";
      orderby = "desc";
      setOrderBy(orderby);
    } else {
      e.target.className = "order asc";
      orderby = "asc";
      setOrderBy(orderby);
    }

    let dataUrl = url.split("=");
    if (dataUrl.length === 2) {
      setUrl(`${dataUrl[0]}=${dataUrl[1]}&order=${orderby}`);
    } else if (dataUrl.length === 4) {
      setUrl(`${dataUrl[0]}=${dataUrl[1]}=${orderby}&search=${dataUrl[3]}`);
    } else setUrl(`${dataUrl[0]}=${dataUrl[1]}=${orderby}`);
  };

  let cont = 0;
  return (
    <>
      {method === "GET" && players.length > 0 && (
        <div
          className="order asc"
          onClick={handleChangeOrder}
          id="order-by"
        ></div>
      )}
      <div className="wrapper">
        {players.map((player) => {
          if (cont === 4) cont = 1;
          else cont++;
          return <PlayerCard key={player.player_id} {...player} img={cont} />;
        })}
      </div>
      {players.length === 0 && <Warning />}
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
