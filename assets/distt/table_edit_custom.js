$(document).ready(function() {
    $('#data_table').Tabledit({
        deleteButton: false,
        editButton: false,
        columns: {
            identifier: [0, 'Key'],
            editable: [
                [1, 'Name, Company, Sector'],
                [2, 'Number of founders'],
                [3, 'Level of Innovation on Joining Dare 2020'],
                [4, 'Current Level of the innovation'],
                [5, 'Pivoted or Preserved'],
                [6, 'Email Address'],
                [7, 'Phone Number'],
                [8, 'Gender'],
                [9, 'Date Joined'],
                [10, 'Moodle Completion'],
                [11, 'Important Links'],
            ]
        },
        hideIdentifier: true,
        url: ""
    });
});

$(document).ready(function() {
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
        url: "test"
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
        url: ""
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
        url: ""
    });
});