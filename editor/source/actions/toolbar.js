import { ACTION, utilAction } from './index';

export function saveContent() {
    return dispatch => {
        let state = globalStore.getState();
        let currentId = state.editor.get('id');
        let currentContent = state.editor.get('current');
        dispatch(utilAction.showTip('load'));
        utilAction.apiRequest('PUT', `articles/${currentId}`, {
            content: currentContent
        }).then(function(data) {
            dispatch(utilAction.showTip('auto', '保存成功'));
        });
    };
}
