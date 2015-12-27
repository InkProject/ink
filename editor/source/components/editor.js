import React from 'react';
import Component from './index';
import classNames from 'classnames';
import ace from 'brace';
import 'brace/mode/markdown';
import 'brace/theme/tomorrow';

import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { listAction, editorAction } from '../actions';

import Header from './header';

class Editor extends Component {
    constructor(props) {
        super(props);
        this.state = {configMode: false};
    }
    switchConfigMode() {
        this.setState({configMode: !this.state.configMode});
    }
    resizeEditor() {
        let width = window.innerWidth ||
            document.documentElement.clientWidth ||
            document.body.clientWidth;
        if (width > 700) {
            this.editor.renderer.setPadding((width - 700) / 2);
            this.configEditor.renderer.setPadding((width - 700) / 2);
        }
    }
    // shouldComponentUpdate(nextProps, nextStates) {
    //     alert(nextProps.params.id != this.props.params.id)
    //     return nextProps.params.id != this.props.params.id;
    // }
    componentDidUpdate(prevProps, prevState) {
        if (this.props.content != prevProps.content) {
            this.editor.setValue(this.props.content || '', -1);
        }
        if (this.props.config != prevProps.config) {
            this.configEditor.setValue(this.props.config || '', -1);
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
        editor.on('focus', () =>
            this.props.listAction.hideList()
        );
    }
    componentDidMount () {
        // init content editor
        this.editor = ace.edit('content-editor');
        this.setEditorStyle(this.editor);
        // init config editor
        this.configEditor = ace.edit('config-editor');
        this.setEditorStyle(this.configEditor);
        this.configEditor.on('input', () => {
            this.props.editorAction.setHeader(this.configEditor.getValue())
        });
        // resize by window size
        this.resizeEditor();
        window.addEventListener('resize', this.resizeEditor.bind(this));
    }
    render() {
        return (
            <div className="editor-wrap">
                <Header title={this.props.title} tags={this.props.tags} edit={this.state.configMode} onClick={() => this.switchConfigMode()} />
                <div className={classNames({hide: this.state.configMode})}><div id="content-editor"></div></div>
                <div className={classNames({hide: !this.state.configMode})}><div id="config-editor"></div></div>
            </div>
        );
    }
}

export default connect(function(state) {
    return {
        title: state.editor.get('title'),
        tags: state.editor.get('tags'),
        config: state.editor.get('config'),
        content: state.editor.get('content')
    }
}, function(dispatch) {
    return {
        listAction: bindActionCreators(listAction, dispatch),
        editorAction: bindActionCreators(editorAction, dispatch)
    };
})(Editor);
