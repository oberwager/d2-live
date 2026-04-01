<div align="center">
<h1>D2-Live</h1>
<h2>Generate Live Diagram URLs with D2</h2>

[![Publish Status](https://github.com/oberwager/d2-live/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/oberwager/d2-live/actions/workflows/docker-publish.yml)
[![Build Status](https://github.com/oberwager/d2-live/actions/workflows/build.yml/badge.svg)](https://github.com/oberwager/d2-live/actions/workflows/build.yml)
[![Uptime](https://uptime.lucas.tools/api/badge/2/uptime/240)](https://uptime.lucas.tools/)
[![License](https://img.shields.io/badge/License-MIT-orange.svg)](https://github.com/oberwager/d2-live/blob/master/LICENSE)
[![Coverage Status](https://coveralls.io/repos/github/oberwager/d2-live/badge.svg?branch=master)](https://coveralls.io/github/oberwager/d2-live?branch=master)
[![Buy Me A Coffee](https://img.shields.io/badge/Buy%20Me%20A%20Coffee-yellow?style=flat-square&logo=buy-me-a-coffee&logoColor=white)](https://www.buymeacoffee.com/lucasobe)
</div>

## Aboout
D2-Live is a dynamic tool for creating live, embeddable diagrams using the D2 language.

## Key Features and Benefits
- **Effortless Diagram Generation**: Convert text to diagrams instantly.
- **Live URLs**: Share and embed your diagrams anywhere with live URLs.
- **SVG Output**: High-quality, scalable vector graphics ensure your diagrams look sharp at any size.

## How to Use D2-Live
1. **Create a Diagram**: Start with a text-based diagram in D2 language. Example:
   ```
   x: That's got to be the best pirate I've ever seen
   x -> y
   y: So it would seem...
   y -> z: Jack Sparrow's theme starts
   z: Fade into horizon
   ```
2. **Use D2-Playground**: Paste your diagram into [d2-playground](https://play.d2lang.com/).
3. **Generate Live URL**: Change the URL domain to `d2.lucas.tools`. For example, `https://d2.lucas.tools/?script=<encoded_string>` gives you an SVG of your diagram.
4. **Embed and Share**: The SVG diagram can be embedded and shared anywhere.

![Diagram Example](https://d2.lucas.tools/?script=FMwxDsIwDAXQPaf4Wyd6gAyMSLCWC7jUIhE0rmzTNjk9yvykd0Y8E_lgeIvDBTPDE2Nmc2xZyRn3YWfwzgpjLuHE5YoaasQkyI5Dft-l0zqOY6hdW8SDXh9MG6nKMVg_V4Y5qVtoETdaGLm4IInmJiX8AwAA__8%3D&)

You can use any of the [D2-Themes](https://d2lang.com/tour/themes) to customize the look of your diagram, just add `&theme=<theme_id>` to the URL.

## Contributing to D2-Live
We welcome contributions! Please refer to our [Contribution Guidelines](https://github.com/oberwager/d2-live/CONTRIBUTING.md) for details on how to submit changes, coding standards, and testing procedures.

## License
D2-Live is MIT licensed. For more information, please refer to the [LICENSE](https://github.com/oberwager/d2-live/blob/master/LICENSE) file.
