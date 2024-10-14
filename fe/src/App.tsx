import "./App.css";
import { Routes, Route } from "react-router-dom";
import TopPage from "./pages/TopPage";
import TaskManagementPage from "./pages/TaskManagementPage";
import LogPage from "./pages/LogPage";

const App = () => {
  return (
    <div className="App">
      <Routes>
        <Route path="/" element={<TopPage />} />
        <Route path="/tm" element={<TaskManagementPage />} />
        <Route path="/log" element={<LogPage />} />
      </Routes>
    </div>
  );
};

export default App;
