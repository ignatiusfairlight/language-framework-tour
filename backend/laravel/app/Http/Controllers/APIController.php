<?php

namespace App\Http\Controllers;

use Illuminate\Http\JsonResponse;

class APIController
{
    // Return success
    ## Apparently Claude says I should use mixed instead of array of $data, stupid clanker
    protected function success(mixed $data = [], string $message = 'success', int $code = 200): JsonResponse
    {
        return response()->json([
            'success' => true,
            'message' => $message,
            'data' => $data,
        ], $code);
    }

    // Return fail
    protected function fail(array $error = [], string $message = 'Error', int $code = 400): JsonResponse
    {
        return response()->json([
            'success' => false,
            'message' => $message,
            'error' => $error,
        ], $code);
    }
}
