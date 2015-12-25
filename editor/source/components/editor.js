import React from 'react';
import Component from './index';
import ace from 'brace';
import 'brace/mode/markdown';
import 'brace/theme/tomorrow';

import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { listAction, editorAction } from '../actions';

class Editor extends Component {
    resizeEditor () {
        let width = window.innerWidth ||
            document.documentElement.clientWidth ||
            document.body.clientWidth;
        if (width > 700) {
            this.editor.renderer.setPadding((width - 700) / 2);
        }
    }
    componentDidUpdate () {
        this.editor.setValue(this.props.content || '', -1);
    }
    componentDidMount () {
        let editor = ace.edit('editor');
        editor.setOptions({
            scrollPastEnd: true,
            showGutter: false,
            wrap: true,
            theme: 'ace/theme/tomorrow',
            mode: 'ace/mode/markdown',
            showPrintMargin: false,
            fontSize: '14px',
            fontFamily: "Menlo, Consolas, 'source-code-pro', 'DejaVu Sans Mono', Monaco, 'Ubuntu Mono', 'Courier New', Courier, 'Microsoft Yahei', 'Hiragino Sans GB', 'WenQuanYi Micro Hei', monospace",
            hScrollBarAlwaysVisible: false,
            selectionStyle: "line",
            displayIndentGuides: false
        });
        editor.renderer.setScrollMargin(200, 200);
        editor.container.style.lineHeight = 2;
        editor.on('input', () => {
            this.props.editorAction.setHeader(this.editor.getValue())
        });
        editor.on('focus', () =>
            this.props.listAction.hideList()
        );
        this.editor = editor;
        this.resizeEditor();
        window.addEventListener('resize', this.resizeEditor.bind(this));
    }
    render() {
        const id = this.props.params.id;
        if (id) {
            this.props.listAction.openArticle(id);
        }
        return (
            <div id="editor"></div>
        );
    }
}

export default connect(function(state) {
    return {
        content: state.editor.get('content')
    }
}, function(dispatch) {
    return {
        listAction: bindActionCreators(listAction, dispatch),
        editorAction: bindActionCreators(editorAction, dispatch)
    };
})(Editor);
