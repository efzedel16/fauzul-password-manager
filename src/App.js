import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import SignUp from "./components/pages/SignUp";
import SignIn from "./components/pages/SignIn";
import Main from "./components/pages/Main";
import PassApp from "./components/pages/PassApp";

const App = () => {
  return (
    <Router>
      <Switch>
        <Route component={SignUp} exact path="/signup" />
        <Route component={SignIn} exact path="/signin" />
        <Route component={PassApp} exact path="/pass-app" />
        <Route component={Main} exact path="/" />
      </Switch>
    </Router>
  );
};

export default App;