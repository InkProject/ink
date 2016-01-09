import * as listAction from './list';
import * as editorAction from './editor';
import * as toolbarAction from './toolbar';
import * as utilAction from './util';

const apiURL = 'http://localhost:8000';

const ACTION = {
    SHOW_LIST: 'SHOW_LIST',
    HIDE_LIST: 'HIDE_LIST',
    SHOW_LOADING: 'SHOW_LOADING',
    REFRESH_LIST: 'REFRESH_LIST',

    SET_HEADER: 'SET_HEADER',
    SET_CONTENT: 'SET_CONTENT',

    SAVE_CONTENT: 'SAVE_CONTENT',

    SHOW_TIP: 'SHOW_TIP'
};

export {
    ACTION,
    apiURL,
    listAction,
    editorAction,
    toolbarAction,
    utilAction
};
