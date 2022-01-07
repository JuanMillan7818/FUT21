import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { App } from "../components/App";
import { ContentMain } from "../components/content/ContentMain";
import { PlayerExpand } from "../components/player/PlayerExpand";


export const AppRouter = () => {  
  return (
    <BrowserRouter>    
      <App />
      <Routes>
        <Route path="/" element={<ContentMain />} />
        <Route path="/search" element={<PlayerExpand />} />        
      </Routes>
    </BrowserRouter>
  );
};

