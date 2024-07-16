import {createTheme, rem} from "@mantine/core";

export const theme = createTheme({
    fontFamily: 'Montserrat, sans-serif',
    fontSizes: {
        h1: rem('34px'),
        h2: rem('26px'),
        h3: rem('22px'),
        h4: rem('18px'),
        h5: rem('16px'),
        h6: rem('14px'),
        xl: rem('20px'),
        lg: rem('18px'),
        md: rem('10px'),
        sm: rem('14px'),
        xs: rem('12px'),
    },
    colors: {
        'light':
            [
                "#D6EBF6",
                "#C4E2F2",
                "#B3DAEE",
                "#A3D2EA",
                "#94CBE7",
                "#85C4E3",
                "#77BDE0",
                "#69B6DD",
                "#5CB0DA",
                "#4FAAD7"
            ],
        'dark':
            [
                "#eef4f6",
                "#cddee4",
                "#acc7d2",
                "#8bb1c0",
                "#6a9bae",
                "#518195",
                "#3f6474",
                "#2d4853",
                "#1b2b32",
                "#090e11",
            ],
        'secondaryColor':
            [
                "#E977A3",
                "#E76899",
                "#E55B8F",
                "#E34D87",
                "#E1417E",
                "#DF3576",
                "#DD296E",
                "#D72267",
                "#CC2062",
                "#C21F5D",
            ],
    },
    primaryColor: 'dark',
    colorScheme: 'dark',
    // Primary shade is used in all components to determine which color from theme.colors[color] should be used.
    primaryShade: {light: 5, dark: 5},
    shadows: {
        'normal': ['0px 6px 16px 0px #0000000F']
    },
    radius: {
        'normal': rem('4px')
    },
    defaultRadius: 'sm',
    focusRing: 'never',
    // breakpoints: {
    //     xs: 0,
    //     sm: 480,
    //     md: 768,
    //     lg: 1024,
    //     xl: 1440,
    // }
});