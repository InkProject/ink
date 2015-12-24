import { combineReducers } from 'redux';
import { routeReducer } from 'redux-simple-router';

import list from './list';
import editor from './editor';

export default combineReducers({
    list,
    editor,
    routing: routeReducer
});
