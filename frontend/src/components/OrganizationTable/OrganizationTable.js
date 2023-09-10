import React from "react";
import "./OrganizationTable.css"; // Import the CSS

function OrganizationTable({ results }) {
  return (
    <table className="organization-table">
      <thead>
        <tr className="table-header">
          <th>Name</th>
          <th>Website</th>
          <th>Country</th>
          <th>Description</th>
          {/* Add other fields as required */}
        </tr>
      </thead>
      <tbody>
        {results && results.length > 0 ? results.map((org) => (
          <tr key={org.Organization_ID}>
            <td>{org.Name}</td>
            <td>{org.Website}</td>
            <td>{org.Country}</td>
            <td>{org.Description}</td>
            {/* Add other fields as required */}
          </tr>
        )) : (
          <tr>
            <td colSpan="4">No organizations found</td>
          </tr>
        )}
      </tbody>
    </table>
  );
}

export default OrganizationTable;
