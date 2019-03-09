<?php
include('vendor/rmccue/requests/library/Requests.php');
Requests::register_autoloader();
$headers = array();
$data = array(
    'From' => '+19542405000',
    'To' => '73297#19541235001',
    'Url' => 'https://cloud.restcomm.com/restcomm-rvd/services/apps/AP1c9836a2f3d04e72a74f1e277d232e42/controller'
);
$response = Requests::post('https://ACcc38b3e0fcf3d2654cd73e210f542233ce:13dd0e03b368f49586428e7a5a3sc49a@cloud.restcomm.com/restcomm/2012-04-24/Accounts/ACcc38b3e0fcf3d2654cd73e210f5477ce/Calls.json', $headers, $data);
