import { combineReducers } from 'redux';
import { routeReducer } from 'redux-simple-router';

import menu from './menu';
import editor from './editor';

export default combineReducers({
    menu,
    editor,
    routing: routeReducer
});
