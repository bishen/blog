buttons: [
        // .... ,
        "insertCode" //custom button
    ],


    customButtons: {
        insertCode: {
            title: 'Insert Code',
            icon: {
                type: 'font',
                value: 'fa fa-dollar' // Font Awesome icon class fa fa-*
            },
            callback: function(editor) {
                editor.saveSelection();

                var codeModal = $("<div>").addClass("froala-modal").appendTo("body");
                var wrapper = $("<div>").addClass("f-modal-wrapper").appendTo(codeModal);
                $("<h4>").append('<span data-text="true">Insert Code</span>')
                    .append($('<i class="fa fa-times" title="Cancel">')
                        .click(function() {
                            codeModal.remove();
                        }))
                    .appendTo(wrapper);

                var dialog = "<textarea id='code_area' style='height: 211px; width: 538px;' /><br/><label>Language:</label><select id='code_lang'><option>CSharp</option><option>VB</option><option>JScript</option><option>Sql</option><option>XML</option><option>CSS</option><option>Java</option><option>Delphi</option></select> <input type='button' name='insert' id='insert_btn' value='Insert' /><br/>";
                $(dialog).appendTo(wrapper);

                $("#code_area").text(editor.text());

                if (!editor.selectionInEditor()) {
                    editor.$element.focus();
                }

                $('#insert_btn').click(function() {
                    var lang = $("#code_lang").val();
                    var code = $("#code_area").val();
                    code = code.replace(/\s+$/, ""); // rtrim
                    code = $('<span/>').text(code).html(); // encode        

                    var htmlCode = "<pre language='" + lang + "' name='code'>" + code + "</pre></div>";
                    var codeBlock = "<div align='left' dir='ltr'>" + htmlCode + "</div><br/>";

                    editor.restoreSelection();
                    editor.insertHTML(codeBlock);
                    editor.saveUndoStep();

                    codeModal.remove();
                });
            }
        }
    }