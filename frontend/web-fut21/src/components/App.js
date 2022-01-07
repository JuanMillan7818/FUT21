import React from "react";
import "../styles/style.css";
import { Search } from "./search/Search";

export const App = () => {
  return (
    <>
      <div className="container animate__animated animate__fadeIn animate__fast">
        <div className="layer">
          <div className="parallelo"></div>
          <div className="imgBx"></div>
          <div className="content">
            <div className="title-main">
              <h1>FIFA 21</h1>
              <h1>Ultimate Team</h1>
            </div>
            <div className="text-main">
              <div className="text-left">
                <p>
                  This FIFA Ultimate Team tool will allow us to find any player
                  of the time and see the members of the existing clubs.
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <Search searchDefault={"player"}/>
    </>
  );
};
