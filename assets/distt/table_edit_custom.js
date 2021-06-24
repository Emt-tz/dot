$(document).ready(function() {
    $('#data_table').Tabledit({
        deleteButton: false,
        editButton: true,
        columns: {
            identifier: [0, 'id'],
            editable: [
                [1, "Detail"],
                [2, "Value"],

            ]
        },
        hideIdentifier: true,
        url: null,
        onAjax: function(action, serialize) {
            console.log(serialize);
            var url = "p/progress?edit=table1&table=Table.&" + serialize;
            fetch(url)
                .then(response => response.json())
                .then(data => {
                    console.log(data.response);
                })
                .catch(err => console.log(err))

        }
    });
});

$(document).ready(function() {
    var url;
    $('#data_table1').Tabledit({
        deleteButton: false,
        editButton: true,
        columns: {
            identifier: [0, 'id'],
            editable: [
                [1, 'Intervention'],
                [2, 'Lead'],
                [3, 'Date'],
                [4, 'Participation'],
                [5, 'Impact'],
                [6, 'Scoring'],
                [7, 'Outcome'],
            ]
        },
        hideIdentifier: true,
        onAjax: function(action, serialize) {
            console.log('onAjax(action, serialize)');
            console.log(action);
            console.log(serialize);
            var url = "p/progress?edit=table&table=Table2.&" + serialize;
            fetch(url)
                .then(response => response.json())
                .then(data => {
                    console.log(data.response);
                })
                .catch(err => console.log(err))

        }


    });

});

$(document).ready(function() {
    $('#data_table2').Tabledit({
        deleteButton: false,
        editButton: true,
        columns: {
            identifier: [0, 'id'],
            editable: [
                [1, 'Intervention'],
                [2, 'Lead'],
                [3, 'Date'],
                [4, 'Participation'],
                [5, 'Impact'],
                [6, 'Scoring'],
                [7, 'Outcome'],
            ]
        },
        hideIdentifier: true,
        onAjax: function(action, serialize) {
            console.log('onAjax(action, serialize)');
            console.log(action);
            console.log(serialize);
            var url = "p/progress?edit=table&table=Table3.&" + serialize;
            fetch(url)
                .then(response => response.json())
                .then(data => {
                    console.log(data.response);
                })
                .catch(err => console.log(err))

        },
        onFail: function(jqXHR, textStatus, errorThrown) {
            console.log('onFail(jqXHR, textStatus, errorThrown)');
            console.log(jqXHR);
            console.log(textStatus);
            console.log(errorThrown);
        },
    });
});

$(document).ready(function() {
    $('#data_table3').Tabledit({
        deleteButton: false,
        editButton: true,
        columns: {
            identifier: [0, 'id'],
            editable: [
                [1, 'Intervention'],
                [2, 'Lead'],
                [3, 'Date'],
                [4, 'Participation'],
                [5, 'Impact'],
                [6, 'Scoring'],
                [7, 'Outcome'],
            ]
        },
        hideIdentifier: true,
        onAjax: function(action, serialize) {
            console.log('onAjax(action, serialize)');
            console.log(action);
            console.log(serialize);
            var url = "p/progress?edit=table&table=Table4.&" + serialize;
            fetch(url)
                .then(response => response.json())
                .then(data => {
                    console.log(data.response);
                })
                .catch(err => console.log(err))

        }
    });
});