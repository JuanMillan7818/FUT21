import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { App } from "../components/App";
import { ContentMain } from "../components/content/ContentMain";


export const AppRouter = () => {  
  return (
    <BrowserRouter>    
      <App />
      <Routes>
        <Route path="/" element={<ContentMain />} />
        <Route path="/search" element={<ContentMain />} />
        <Route path="/team" element={<ContentMain />} />
      </Routes>
    </BrowserRouter>
  );
};

