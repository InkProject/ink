import * as listAction from './list'
import * as editorAction from './editor'
import * as toolbarAction from './toolbar'
import * as menuAction from './menu'
import * as modalAction from './modal'
import * as utilAction from './util'

const apiURL = 'http://localhost:8000'

const ACTION = {
    SHOW_LIST: 'SHOW_LIST',
    HIDE_LIST: 'HIDE_LIST',
    SHOW_LOADING: 'SHOW_LOADING',
    REFRESH_LIST: 'REFRESH_LIST',

    SET_HEADER: 'SET_HEADER',
    SET_CONTENT: 'SET_CONTENT',
    SET_CURRENT: 'SET_CURRENT',

    SAVE_CONTENT: 'SAVE_CONTENT',

    SHOW_TIP: 'SHOW_TIP',

    SHOW_MODAL: 'SHOW_MODAL',
    CREATE_ARTICLE: 'CREATE_ARTICLE'
}

export {
    ACTION,
    apiURL,
    listAction,
    editorAction,
    toolbarAction,
    utilAction,
    menuAction,
    modalAction
}
