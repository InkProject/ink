import React from 'react';
import Component from './index';
import classNames from 'classnames';
import ace from 'brace';
import 'brace/mode/markdown';
import 'brace/theme/tomorrow';
import _ from 'lodash';

import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { listAction, editorAction, utilAction } from '../actions';

import Header from './header';

class Editor extends Component {
    constructor(props) {
        super(props);
        this.state = {configMode: false};
    }
    switchConfigMode() {
        this.setState({configMode: !this.state.configMode});
        if (this.state.configMode) {
            this.contentEditor.focus();
            this.props.utilAction.showTip('auto', '切换至内容');
        } else {
            this.configEditor.focus();
            this.props.utilAction.showTip('auto', '切换至配置');
        }
        this.setEditorStyle(this.contentEditor);
        this.setEditorStyle(this.configEditor);
    }
    resizeEditor() {
        let width = window.innerWidth ||
            document.documentElement.clientWidth ||
            document.body.clientWidth;
        if (width > 700) {
            this.contentEditor.renderer.setPadding((width - 700) / 2);
            this.configEditor.renderer.setPadding((width - 700) / 2);
        }
        this.setEditorStyle(this.contentEditor);
        this.setEditorStyle(this.configEditor);
    }
    componentDidUpdate(prevProps, prevState) {
        if (this.props.editor.get('id') != this.props.params.id) {
            this.props.listAction.openArticle(this.props.params.id);
        }
        if (!prevProps || this.props.editor.get('id') != prevProps.editor.get('id')) {
            this.contentEditor.setValue(this.props.editor.get('content') || '', -1);
            this.configEditor.setValue(this.props.editor.get('config') || '', -1);
            this.setState({configMode: false});
        }
    }
    setEditorStyle(editor) {
        const editorOption = {
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
        };
        editor.setOptions(editorOption);
        editor.renderer.setScrollMargin(200, 200);
        editor.container.style.lineHeight = 2;
        editor.$blockScrolling = Infinity;
        editor.on('focus', () =>
            this.props.listAction.hideList()
        );
    }
    componentDidMount () {
        // init content editor
        this.contentEditor = ace.edit('content-editor');
        this.setEditorStyle(this.contentEditor);
        // init config editor
        this.configEditor = ace.edit('config-editor');
        this.setEditorStyle(this.configEditor);
        this.configEditor.on('input', () => {
            this.props.editorAction.setHeader(this.configEditor.getValue())
            this.onEditorChange();
        });
        this.contentEditor.on('input', () => {
            this.onEditorChange();
        });
        // resize by window size
        this.resizeEditor();
        window.addEventListener('resize', this.resizeEditor.bind(this));
        this.componentDidUpdate();
    }
    onEditorChange () {
        let config = _.trim(this.configEditor.getValue());
        let content = _.trim(this.contentEditor.getValue());
        let current = `${_.trim(config)}\n\n---\n\n${_.trim(content)}`;
        this.props.editorAction.setCurrent(current);
    }
    render() {
        const editor = this.props.editor;
        return (
            <div className="editor-wrap">
                <Header title={editor.get('title')} tags={editor.get('tags')} edit={this.state.configMode} onClick={() => this.switchConfigMode()} />
                <div className={classNames({hide: this.state.configMode})}><div id="content-editor"></div></div>
                <div className={classNames({hide: !this.state.configMode})}><div id="config-editor"></div></div>
            </div>
        );
    }
}

export default connect(function(state) {
    return {
        editor: state.editor
    }
}, function(dispatch) {
    return {
        listAction: bindActionCreators(listAction, dispatch),
        editorAction: bindActionCreators(editorAction, dispatch),
        utilAction: bindActionCreators(utilAction, dispatch)
    };
})(Editor);
