import React, {useState} from "react";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";

import Navbar from "./components/Navbar";
import SideBar from "./components/SideBar";

import Website from "./pages/Website";
import BusinessContext from "./pages/BusinessContext";
import ConceptualModel from "./pages/ConceptualModel";
import UserStory from "./pages/UserStory";

function App() {
  const [docs] = useState([
    { path: "/ruaairline-website", name: "Rùa Airline Website", component: <Website />},
    { path: "/business-context", name: "Business Context", component: <BusinessContext />},
    { path: "/conceptual-model", name: "Conceptual Model", component: <ConceptualModel />},
    { path: "/user-story", name: "User Story", component: <UserStory />}
  ]);
  return (
    <Router>
      <div className="flex">
        <SideBar items={docs} />
        <div className="flex-1">
          <Navbar />
          <div className="p-4">
            <Routes>
              {docs.map((doc, index) => (
                <Route key={index} path={doc.path} element={doc.component} />
              ))}
              <Route path="/" element={<Website />} />
            </Routes>
          </div>
        </div>
      </div>
    </Router>
  );
}

export default App;
