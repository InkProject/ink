import { combineReducers } from 'redux';

import menu from './menu';
import editor from './editor';

export default combineReducers({
    menu,
    editor
});
