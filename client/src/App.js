import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import SignUp from "./components/pages/SignUp";
import SignIn from "./components/pages/SignIn";

function App() {
  return (
    <Router>
      <Switch>
        <Route component={SignUp} exact path="/signup" />
        <Route component={SignIn} exact path="/signin" />
      </Switch>
    </Router>
  );
}

export default App;