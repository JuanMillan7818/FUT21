import React, { useContext, useState } from "react";
import PropTypes from "prop-types";
import { ApiContext } from "../context/ApiContext";

export const Search = ({ searchDefault }) => {
  const { setUrl, setParams} = useContext(ApiContext);

  const [search, setSearch] = useState("");
  const [typeSearch, setTypeSearch] = useState(searchDefault);  

  const handleInputChange = ({ target }) => {
    setSearch(target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(search, typeSearch);        
    if (typeSearch === "player") {      
      setParams({
        method: 'GET',        
      })

      setUrl(
        `${process.env.REACT_APP_URL_PLAYER}page=1&order=asc&search=${search.trim()}`
      );      
    }else {
      const json = {
        Name: search.trim(),
        Page: 1
      }
      setParams({        
        method: 'POST',
        body: JSON.stringify(json)
      })
      setUrl(`${process.env.REACT_APP_URL_TEAM}`)
    }
    setSearch('')
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
