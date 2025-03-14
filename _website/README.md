# LLPkgStore Site
This website provides the download service of `llpkgstore.json` for `llgo get`, and meanwhile provide service that users of llgo can look up for the version mapping of LLPkg.

The `llpkgstore.json` is stored in `/_website/public`

## Installation
To install the necessary dependencies using `corepack` and `yarn`, move to `/_website` and follow these steps:

1. **Enable Corepack** (if not already enabled):
    ```sh
    corepack enable
    ```

2. **Install Dependencies**:
    ```sh
    yarn
    ```

This will install all the required packages listed in the `package.json` file.

## Usage
To start the development server, run:

```sh
yarn dev
```

This will launch the website locally, and you can view it in your browser at `http://localhost:5173`.

## Building for Production
To build the website for production, run:

```sh
yarn build
```

This will generate the static files in the `dist` directory, which can be deployed to your hosting provider.