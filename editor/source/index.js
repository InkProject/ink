import React from 'react';
import ReactDom from 'react-dom';
import { createStore, applyMiddleware } from 'redux';
import { Provider } from 'react-redux';
import thunk from 'redux-thunk';
import rootReducer from './reducers';
import App from './app';
import List from './components/list';
import Editor from './components/editor';
import { Router, Route, IndexRoute } from 'react-router';
import { syncReduxAndRouter } from 'redux-simple-router';
import { createHistory } from 'history';

const createStoreWithMiddleware = applyMiddleware(thunk)(createStore);
const store = createStoreWithMiddleware(rootReducer);

const history = createHistory();
syncReduxAndRouter(history, store)

ReactDom.render(
    <Provider store={store}>
        <Router history={history}>
            <Route path="/" component={App}>
                <IndexRoute components={{List, Editor}}/>
                <Route path="/edit/:id" components={{List, Editor}}/>
            </Route>
        </Router>
    </Provider>,
    document.getElementById('root')
);
