import { createStore, applyMiddleware } from 'redux'
import thunkMiddleware from 'redux-thunk'
import rootReducer from './reducer'
import { syncHistory, routeReducer } from 'react-router-redux'
import history from './history'

const reduxRouterMiddleware = syncHistory(history)
const createStoreWithMiddleware = applyMiddleware(thunkMiddleware, reduxRouterMiddleware)(createStore)
const store = createStoreWithMiddleware(rootReducer)
reduxRouterMiddleware.listenForReplays(store)

if (module.hot) {
    module.hot.accept('./reducer', () => {
        const nextRootReducer = require('./reducer')
        store.replaceReducer(nextRootReducer)
    })
}

window.store = store

export default store
