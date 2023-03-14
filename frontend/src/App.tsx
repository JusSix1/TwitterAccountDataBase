import * as React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import './App.css';
import FullAppBar from "./component/FullAppBar";
import Home_User from "./component/Home_User";
import SignIn_User from "./component/SignIn_User_UI";

export default function App() {
  const [token, setToken] = React.useState<String>("");

  React.useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setToken(token);
    }
  }, []);

  if (!token) {
    return <SignIn_User />;
  }

  function routeList() {
    if(localStorage.getItem("position") == "Admin"){
      // return( // Admin Routes
      // );
    }else{ // User Routes
      return(
        <Routes>
          <Route path="/" element={<Home_User/>} /> {/** home */}
        </Routes>
      );
    }
  }

  return (
  <Router>
    <div>
      <FullAppBar />
      {routeList()}
    </div>
  </Router>
  );
}
