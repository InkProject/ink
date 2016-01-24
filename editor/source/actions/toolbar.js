import { ACTION, utilAction, listAction } from './index';
import history from '../history';

export function saveContent() {
    return dispatch => {
        let state = globalStore.getState();
        let currentId = state.editor.get('id');
        let currentContent = state.editor.get('current');
        dispatch(utilAction.showTip('load'));
        utilAction.apiRequest('PUT', `articles/${currentId}`, {
            content: currentContent
        }).then(function(data) {
            dispatch(listAction.fetchList());
            dispatch(utilAction.showTip('auto', '保存成功'));
        });
    };
}

export function removeArticle() {
    return dispatch => {
        let state = globalStore.getState();
        let currentId = state.editor.get('id');
        dispatch(utilAction.showTip('load'));
        utilAction.apiRequest('DELETE', `articles/${currentId}`).then(function(data) {
            history.replaceState(null, `/`);
            dispatch(listAction.fetchList());
            dispatch(utilAction.showTip('auto', '删除成功'));
        });
    };
}
