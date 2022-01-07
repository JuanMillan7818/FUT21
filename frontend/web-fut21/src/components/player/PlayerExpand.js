import React, { useContext } from "react";
import { ApiContext } from "../context/ApiContext";

export const PlayerExpand = () => {
  const { searchExpand } = useContext(ApiContext);
  return (
    <>
      {Object.keys(searchExpand).length !== 0 && (
        <div className="container-spand animate__animated animate__bounceInLeft animate__slow">
          <div className="card-spand">
            <div className="content-left">
              <div className={"img-spand " + searchExpand.classImg}></div>
            </div>
            <div className="info">
              <h2>
                {searchExpand.firstName} {searchExpand.lastName}
              </h2>
              <p>
                <strong>Position: </strong> {searchExpand.positionFull}
              </p>
              <p>
                <strong>Nation: </strong> {searchExpand.nation}
              </p>
              <p>
                <strong>Club: </strong> {searchExpand.club}{" "}
              </p>
            </div>
          </div>
        </div>
      )}
    </>
  );
};
