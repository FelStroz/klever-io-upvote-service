import React from "react";
import {
    BrowserRouter as Router,
    Switch,
    Route,
} from "react-router-dom";
import InsertForm from "./pages/Insert";
import Home from "./pages/Home";

export default function Routes() {
    return (
        <Router>
            <Switch>
                <Route exact path={'/'}>
                    <Home />
                </Route>
                <Route exact path={'/insert'}>
                    <InsertForm />
                </Route>
                <Route path="/">
                    Não encontrado 404
                </Route>
            </Switch>
        </Router>
    )
}
