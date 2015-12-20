import * as menuAction from './menu';
import * as editorAction from './editor';

const ACTION = {
    SHOW_LIST: 'SHOW_LIST',
    HIDE_LIST: 'HIDE_LIST',
    SHOW_LOADING: 'SHOW_LOADING',
    SELECT_ARTICLE: 'SELECT_ARTICLE',
    REFRESH_LIST: 'REFRESH_LIST',
    SET_CONTENT: 'SET_CONTENT'
};

export {
    ACTION,
    menuAction,
    editorAction
};
