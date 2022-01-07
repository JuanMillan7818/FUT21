import React, { useContext, useState } from "react";
import PropTypes from "prop-types";
import { ApiContext } from "../context/ApiContext";
import { useLocation, useNavigate } from "react-router-dom";
import queryString from 'query-string'

export const Search = ({ searchDefault }) => {
  const { url, setUrl } = useContext(ApiContext);

  const [search, setSearch] = useState("");
  const [typeSearch, setTypeSearch] = useState(searchDefault);

  const location = useLocation();
  const navigate = useNavigate();
  

  const handleInputChange = ({ target }) => {
    setSearch(target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(search, typeSearch);
  };

  const handleSelectChange = ({ target }) => {
    setTypeSearch(target.value);
  };

  return (
    <>
      <div className="content-second">
        <h2 className="title-second">Revolution Sports</h2>
        <div className="box">
          <h3>Search for player or club</h3>
          <form onSubmit={handleSubmit}>
            <input
              type="text"
              placeholder="Search..."
              value={search}
              onChange={handleInputChange}
              autoComplete="off"
            />
            <select value={typeSearch} onChange={handleSelectChange}>
              <option value="player">Player</option>
              <option value="team">Team</option>
            </select>
            <input type="submit" value="Search" />
          </form>
        </div>
      </div>
    </>
  );
};

Search.propTypes = {
  searchDefault: PropTypes.string.isRequired,
};
