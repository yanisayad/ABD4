<?php

use EscapeGame\Database;

require_once 'lib/autoload.php';

$faker = Faker\Factory::create();
$escapegame = new Database();

if (empty($argv[1]) == true) {
    $timer = 1000000;
} else {
    $timer = $argv[1]*1000;
}


while (true) {
    // usleep($timer);
    $firstAge = $escapegame->getAge();
    $nbTickets = $escapegame->getNbTickets();
    $firstCivility = $escapegame->getGender(false);
    $firstNom = $faker->firstName;
    $firstPrenom = $faker->firstName;
    $firstEmail = $firstNom . '.' . $firstPrenom . '@' . $escapegame->getSuffixEmail();
    $firstPersonType = $escapegame->getPersonType($firstAge);

    $result['Acheteur']['Civilite'] = $firstCivility;
    $result['Acheteur']['Nom'] = $firstNom;
    $result['Acheteur']['Prenom'] = $firstPrenom;
    $result['Acheteur']['Age'] = $firstAge;
    $result['Acheteur']['Email'] = strtolower($firstEmail);
    $result['Game']['Nom'] = $escapegame->getEscapeGameName();
    $result['Game']['Jour'] = $escapegame->getReservationDate();
    $result['Game']['Horaire'] = $escapegame->getReservationHour();
    $result['Game']['VR'] = $escapegame->useVirtualReality();
    for ($i = 0; $i < $nbTickets; $i++) {
        if ($i == 0) {
            $result['Reservation'][$i]['Spectateur']['Civilite'] = $firstCivility;
            $result['Reservation'][$i]['Spectateur']['Nom'] = $firstNom;
            $result['Reservation'][$i]['Spectateur']['Prenom'] = $firstPrenom;
            $result['Reservation'][$i]['Spectateur']['Age'] = $firstAge;
            $result['Reservation'][$i]['Tarif'] = $firstPersonType;
        } else {
            $otherAge = $escapegame->getAge();
            $otherCivility = $escapegame->getGender(false);
            $otherNom = $faker->firstName;
            $otherPrenom = $faker->firstName;
            $otherPersonType = $escapegame->getPersonType($otherAge);

            $result['Reservation'][$i]['Spectateur']['Civilite'] = $otherCivility;
            $result['Reservation'][$i]['Spectateur']['Nom'] = $otherNom;
            $result['Reservation'][$i]['Spectateur']['Prenom'] = $otherPrenom;
            $result['Reservation'][$i]['Spectateur']['Age'] = $otherAge;
            $result['Reservation'][$i]['Tarif'] = $otherPersonType;
        }
    }

    $json_result = json_encode($result);
    $options = [
        CURLOPT_CUSTOMREQUEST => "POST",
        CURLOPT_POSTFIELDS => $json_result,
        CURLOPT_RETURNTRANSFER => true,
        CURLOPT_FOLLOWLOCATION => true,
        CURLOPT_POSTREDIR => 3,
        CURLOPT_HTTPHEADER => [
            'Content-Type: application/json',
            'Content-Length: ' . strlen($json_result)
        ]
    ];
    $addr = 'http://127.0.0.1:8000/transaction';
    $ch = curl_init($addr);
    echo "\n\n";
    echo $addr;
    echo "\n\n";
    curl_setopt_array($ch,$options);
    echo(json_encode($result));
    echo "\n\n";

    $response = curl_exec($ch);
    echo "Response: ".$response;
    echo "\n\n";
    curl_close($ch);
    echo json_encode($response);
    echo "\n\n";
}
