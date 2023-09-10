import React, { useState } from "react";
import "./App.css";
import SearchForm from "./components/SearchForm/SearchForm";
import OrganizationTable from "./components/OrganizationTable/OrganizationTable";

function App() {
  const [name, setName] = useState("");
  const [country, setCountry] = useState("");
  const [results, setResults] = useState([]);

  const fetchAllOrganizations = async () => {
    console.log("Fetching all organizations..."); // Debug

    try {
      const response = await fetch(`http://localhost:8080/organizations`);

      if (!response.ok) {
        throw new Error("Network response was not ok");
      }

      const data = await response.json();
      setResults(data);
    } catch (error) {
      console.error("Fetch error:", error);
    }
  };

  const searchOrganizations = async () => {
    const response = await fetch(
      `http://localhost:8080/search?name=${name}&country=${country}`
    );
    const data = await response.json();
    setResults(data);
  };

  const updateData = async () => {
    await fetch("http://localhost:8080/update", { method: "POST" });
    alert("Data updated!");
  };

  return (
    <div className="App">
      <h1>Organization Search</h1>
      <SearchForm
        name={name}
        country={country}
        onNameChange={(e) => setName(e.target.value)}
        onCountryChange={(e) => setCountry(e.target.value)}
        onSearchClick={searchOrganizations}
        onUpdateClick={updateData}
      />
      <button onClick={fetchAllOrganizations}>Fetch All Organizations</button>
      <OrganizationTable results={results} />
    </div>
  );
}

export default App;
