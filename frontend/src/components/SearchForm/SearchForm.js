import React from 'react';
import './SearchForm.css'; // Import the CSS file

function SearchForm({ name, country, onNameChange, onCountryChange, onSearchClick, onUpdateClick }) {
  return (
    <div className="search-form"> {/* Add the class name */}
      <input className="search-input" placeholder="Name" value={name} onChange={onNameChange} />
      <input className="search-input" placeholder="Country" value={country} onChange={onCountryChange} />
      <button className="search-button" onClick={onSearchClick}>Search</button>
      <button className="update-button" onClick={onUpdateClick}>Update Data</button>
    </div>
  );
}

export default SearchForm;
