import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import SignUpPage from "./pages/SignUpPage";

const App = () => {
  return (
    <Router>
      <Switch>
        <Route component={SignUpPage} exact path="/user/signup" />
      </Switch>
    </Router>
  );
};

export default App;