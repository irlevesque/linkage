async function apiCall(method, endpoint, body = null) {
  try {
    const options = {
      method: method,
      headers: {
        "Content-Type": "application/json",
      },
    };
    if (body) {
      options.body = JSON.stringify(body);
    }

    const response = await fetch(endpoint, options);
    const data = await response.json();
    if (!response.ok) {
      throw new Error(data.error || `API error: ${response.status}`);
    }
    return data;
  } catch (error) {
    console.error("API Call Error:", error);
    return { error: error.message };
  }
}

function displayOutput(elementId, data) {
  const outputElement = document.getElementById(elementId);
  if (data.error) {
    outputElement.style.color = "red";
    outputElement.textContent = `Error: ${data.error}`;
  } else {
    outputElement.style.color = "black";
    outputElement.textContent = JSON.stringify(data, null, 2);
  }
}

// Links Management Functions
async function getLinks() {
  const data = await apiCall("GET", "/api/links");
  const linksTableBody = document.getElementById("linksTableBody");

  // Clear existing rows
  linksTableBody.innerHTML = "";

  if (data.error) {
    const errorRow = linksTableBody.insertRow();
    const errorCell = errorRow.insertCell();
    errorCell.colSpan = 4; // Span across all columns (Stub, URL, Description and Actions)
    errorCell.style.color = "red";
    errorCell.textContent = `Error: ${data.error}`;
    return;
  }

  if (Array.isArray(data) && data.length > 0) {
    data.forEach((link) => {
      const row = linksTableBody.insertRow();
      const stubCell = row.insertCell();
      const urlCell = row.insertCell();
      const descriptionCell = row.insertCell();
      const deleteCell = row.insertCell();
      stubCell.textContent = link.stub;
      urlCell.textContent = link.url;
      descriptionCell.textContent = link.description;

      // Add delete button
      const deleteButton = document.createElement("button");
      deleteButton.textContent = "❌";
      deleteButton.title = `Delete "${link.stub}" Linkage`;
      deleteButton.onclick = () => deleteLink(link.stub);
      deleteCell.appendChild(deleteButton);
    });
  } else {
    const noLinksRow = linksTableBody.insertRow();
    const noLinksCell = noLinksRow.insertCell();
    noLinksCell.colSpan = 4; // Span across all columns
    noLinksCell.textContent = "No links found.";
  }
}

// Search Engines Management Functions
async function getSearchEngines() {
  const data = await apiCall("GET", "/api/search");
  const searchEnginesTableBody = document.getElementById(
    "searchEnginesTableBody",
  );

  // Clear existing rows
  searchEnginesTableBody.innerHTML = "";

  if (data.error) {
    const errorRow = searchEnginesTableBody.insertRow();
    const errorCell = errorRow.insertCell();
    errorCell.colSpan = 4; // Span across all columns (Name, Query URL, Default and Delete)
    errorCell.style.color = "red";
    errorCell.textContent = `Error: ${data.error}`;
    return;
  }

  if (Array.isArray(data) && data.length > 0) {
    data.forEach((engine) => {
      const row = searchEnginesTableBody.insertRow();
      const nameCell = row.insertCell();
      const queryCell = row.insertCell();
      const defaultCell = row.insertCell();
      const deleteCell = row.insertCell();

      nameCell.textContent = engine.name;
      queryCell.textContent = engine.query_url;

      // Add set default button
      if (!engine.default) {
        const defaultButton = document.createElement("button");
        defaultButton.textContent = "▢";
        defaultButton.title = `Set "${engine.name}" as default search engine`;
        defaultButton.onclick = () => setDefaultSearchEngine(engine.name);
        defaultCell.appendChild(defaultButton);
      } else {
        const defaultButton = document.createElement("button");
        defaultButton.textContent = "✅";
        defaultButton.title = `"${engine.name}" is the current default search engine`;
        defaultCell.appendChild(defaultButton);
      }

      // Add delete button
      const deleteButton = document.createElement("button");
      deleteButton.textContent = "❌";
      deleteButton.title = `Delete "${engine.name}" search engine`;
      deleteButton.onclick = () => deleteSearchEngine(engine.name);
      deleteCell.appendChild(deleteButton);
    });
  } else {
    const noEnginesRow = searchEnginesTableBody.insertRow();
    const noEnginesCell = noEnginesRow.insertCell();
    noEnginesCell.colSpan = 4; // Span across all columns
    noEnginesCell.textContent = "No search engines found.";
  }
}

// Delete Link Function
async function deleteLink(stub) {
  if (
    !confirm(`Are you sure you want to delete the linkage with stub '${stub}'?`)
  ) {
    return;
  }

  const data = await apiCall("DELETE", `/api/link/${stub}`);
  if (data.error) {
    alert(`Error deleting link: ${data.error}`);
  } else {
    alert(`Link '${stub}' deleted successfully`);
    getLinks(); // Refresh the links list
  }
}

// Add Link Function
async function addLink(event) {
  event.preventDefault();

  const stub = document.getElementById("stubInput").value.trim();
  const url = document.getElementById("urlInput").value.trim();
  const description = document.getElementById("descriptionInput").value.trim();

  if (!stub || !url) {
    alert("Stub and URL are required");
    return;
  }

  const newLink = {
    stub: stub,
    url: url,
    description: description || undefined,
  };

  const data = await apiCall("POST", `/api/link/${stub}`, newLink);
  if (data.error) {
    alert(`Error adding link: ${data.error}`);
  } else {
    alert(`Link '${stub}' added successfully`);
    document.getElementById("addLinkForm").reset();
    getLinks(); // Refresh the links list
  }
}

// Delete Search Engine Function
async function deleteSearchEngine(name) {
  if (
    !confirm(`Are you sure you want to delete the search engine '${name}'?`)
  ) {
    return;
  }

  const data = await apiCall("DELETE", `/api/search/${name}`);
  if (data.error) {
    alert(`Error deleting search engine: ${data.error}`);
  } else {
    alert(`Search engine '${name}' deleted successfully`);
    getSearchEngines(); // Refresh the search engines list
  }
}

// Set Default Search Engine Function
async function setDefaultSearchEngine(name) {
  const data = await apiCall("POST", `/api/search/default/${name}`);
  if (data.error) {
    alert(`Error setting default search engine: ${data.error}`);
  } else {
    // alert(`Search engine '${name}' set as default successfully`);
    getSearchEngines(); // Refresh the search engines list
  }
}

// Add Search Engine Function
async function addSearchEngine(event) {
  event.preventDefault();

  const name = document.getElementById("searchEngineNameInput").value.trim();
  const queryUrl = document
    .getElementById("searchEngineQueryUrlInput")
    .value.trim();

  if (!name || !queryUrl) {
    alert("Name and Query URL are required");
    return;
  }

  const newSearchEngine = {
    name: name,
    query_url: queryUrl,
  };

  const data = await apiCall("POST", `/api/search`, newSearchEngine);
  if (data.error) {
    alert(`Error adding search engine: ${data.error}`);
  } else {
    alert(`Search engine '${name}' added successfully`);
    document.getElementById("addSearchEngineForm").reset();
    getSearchEngines(); // Refresh the search engines list
  }
}

document.addEventListener("DOMContentLoaded", () => {
  getLinks();
  getSearchEngines();

  // Add form submission handlers
  document.getElementById("addLinkForm").addEventListener("submit", addLink);
  document
    .getElementById("addSearchEngineForm")
    .addEventListener("submit", addSearchEngine);
});
