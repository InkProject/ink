import React from 'react'
import ReactDom from 'react-dom'
import { Provider } from 'react-redux'
import { Router, Route, IndexRoute } from 'react-router'

import store from './store'
import history from './store/history'

import App from './app'
import List from './components/list'
import Editor from './components/editor'
import Welcome from './components/welcome'

ReactDom.render(
    <Provider store={store}>
        <Router history={history}>
            <Route path="/" component={App}>
                <IndexRoute components={{Welcome}}/>
                <Route path="/edit/:id" components={{Editor}}/>
            </Route>
        </Router>
    </Provider>,
    document.getElementById('root')
)
