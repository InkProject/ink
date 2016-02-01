import { createStore, applyMiddleware } from 'redux'
import thunk from 'redux-thunk'
import rootReducer from './reducer'
import { syncReduxAndRouter } from 'redux-simple-router'
import history from './history'
const createStoreWithMiddleware = applyMiddleware(thunk)(createStore)
const store = createStoreWithMiddleware(rootReducer)

if (module.hot) {
    module.hot.accept('./reducer', () => {
        const nextRootReducer = require('./reducer')
        store.replaceReducer(nextRootReducer)
    })
}

syncReduxAndRouter(history, store)

window.globalStore = store

export default store
