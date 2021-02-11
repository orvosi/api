# Endpoints

## `POST /sign-in`

### Authentication

Bearer token

### Request Body

None

### Request Parameters

None

### Success Response

```
{
    "data": null,
    "meta": {}
}
```

### Error Response

```
{
    "errors": [
        {
            "code": string,
            "message": string
        }
    ],
    "meta": null
}
```

## `POST /medical-records`

### Authentication

Bearer token

### Request Body

```
{
    "symptom": string,
    "diagnosis": string,
    "therapy": string
}
```

### Request Parameters

None

### Success Response

```
{
    "data": null,
    "meta": {}
}
```

### Error Response

```
{
    "errors": [
        {
            "code": string,
            "message": string
        }
    ],
    "meta": null
}
```

## `GET /medical-records`

### Authentication

Bearer token

### Request Body

None

### Request Parameters

None

### Success Response

```
{
    "data": [
        {
            "id": string,
            "symptom": string,
            "diagnosis": string,
            "therapy": string,
            "result": string,
            "created_by": string,
            "created_at": time in string,
            "updated_by": string,
            "updated_at": time in string
        }
    ],
    "meta": {}
}
```

### Error Response

```
{
    "errors": [
        {
            "code": string,
            "message": string
        }
    ],
    "meta": null
}
```

## `GET /medical-records/:id`

### Authentication

Bearer token

### Request Body

None

### Request Parameters

- id: string

### Success Response

```
{
    "data": {
        "id": string,
        "symptom": string,
        "diagnosis": string,
        "therapy": string,
        "result": string,
        "created_by": string,
        "created_at": time in string,
        "updated_by": string,
        "updated_at": time in string
    },
    "meta": {}
}
```

### Error Response

```
{
    "errors": [
        {
            "code": string,
            "message": string
        }
    ],
    "meta": null
}
```

## `PUT /medical-records/:id`

### Authentication

Bearer token

### Request Body

```
{
    "symptom": string,
    "diagnosis": string,
    "therapy": string,
    "result": string
}
```

### Request Parameters

None

### Success Response

```
{
    "data": null,
    "meta": {}
}
```

### Error Response

```
{
    "errors": [
        {
            "code": string,
            "message": string
        }
    ],
    "meta": null
}
```