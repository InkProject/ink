import { combineReducers } from 'redux'
import { routeReducer } from 'react-router-redux'

import list from '../components/list/reducer'
import editor from '../components/editor/reducer'
import util from '../components/util/reducer'
import menu from '../components/menu/reducer'

export default combineReducers({
    list,
    editor,
    util,
    menu,
    routing: routeReducer
})
