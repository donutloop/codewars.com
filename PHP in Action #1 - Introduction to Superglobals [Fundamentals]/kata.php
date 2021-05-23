function double_global_num() {
   $GLOBALS['num'] = $GLOBALS['num'] * 2;
}

function set_query($query) {
    $_GET['query'] = $query;
}

function set_email($email) {
   $_POST['email'] = $email;
}