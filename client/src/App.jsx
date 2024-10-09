import React from "react";
import { Route, Routes } from "react-router-dom";
import About from "./components/AboutApp";
import TodoApp from "./TodoApp";
import AppNavbar from "./components/AppNavbar";

const App = () => {
  return (
    <>
      <AppNavbar />
      <Routes>
        <Route path="/" element={<TodoApp />} />
        <Route path="/about" element={<About />} />
      </Routes>
    </>
  );
};

export default App;
