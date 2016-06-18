/**
 * @license Copyright (c) 2003-2015, CKSource - Frederico Knabben. All rights reserved.
 * For licensing, see LICENSE.md or http://ckeditor.com/license
 */

CKEDITOR.editorConfig = function(config) {
    // Define changes to default configuration here. For example:
    // config.language = 'fr';
    config.uiColor = '#F7F7F7';
    config.height = 500;
    config.extraPlugins = 'codesnippet';
    config.codeSnippet_theme = 'zenburn';
    config.codeSnippet_languages = {
        html: 'HTML',
        go: 'GO',
        css: 'Css',
        php: 'PHP',
        javascript: 'JavaScript',
        apache: 'Apache',
        bash: 'Bash',
        xml: 'Xml',
        json: 'Json',
        coffeescript: 'CoffeeScript',
        java: 'Java',
        cpp: 'c++',
        sql: 'SQL',
        ini: 'INI',
        python: 'Python',
        makefile: 'MakeFile'
    };
};