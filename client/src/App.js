import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import SignUp from "./components/pages/SignUp";
import SignIn from "./components/pages/SignIn";
import Main from "./components/pages/Main";

function App() {
  return (
    <Router>
      <Switch>
        <Route component={SignUp} exact path="/signup" />
        <Route component={SignIn} exact path="/signin" />
        <Route component={Main} exact path="/" />
      </Switch>
    </Router>
  );
}

export default App;