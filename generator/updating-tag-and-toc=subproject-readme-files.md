# Updating TOC TAG and TOC SubProject README Files

This document outlines the process for updating the README files located in the [tags](https://github.com/cncf/toc/tree/main/tags) and [toc_subprojects](https://github.com/cncf/toc/tree/main/toc_subprojects) directories. These README files are automatically generated and updated based on the content of the [`tags.yaml`](https://github.com/cncf/toc/blob/main/tags.yaml) file.

## **The Single Source of Truth: `tags.yaml`**

The `tags.yaml` file serves as the central source of information for generating and updating the README files for both TAGs and TOC SubProjects. To modify the content of these READMEs, you **must** update the `tags.yaml` file directly.

## **Who can Update:**

Anyone can create a pull request in the `cncf/toc` repository, only TOC members can review and merge changes to the `tags.yaml` file.

## **How to Update:**

1.  **Locate the `tags.yaml` file:** This file is located at the root of the `cncf/toc` repository: [https://github.com/cncf/toc/blob/main/tags.yaml](https://github.com/cncf/toc/blob/main/tags.yaml).

2.  **Make your changes:** Carefully edit the `tags.yaml` file to reflect the desired updates for the TAG or subproject README files. Ensure the YAML syntax is correct.

3.  **Create a Pull Request (PR):**
    * **Important:** A PR that aims to update the TAG and subproject README files **must only contain changes to the `tags.yaml` file.** Any other changes in the PR will not trigger the automatic README updates.

## **Automation in Action:**

The `github-project-automation` and `cncf-automation-bot` handle the automatic updates of the README files based on your changes to `tags.yaml`:

* **Updates from a Branch within the `cncf/toc` Repository:**
    * When you create a PR from a branch within the `cncf/toc` repository that modifies only the `tags.yaml` file, the `github-project-automation` will automatically:
        * Generate the necessary updates to the relevant README files in the `tags` and `toc_subprojects` directories.
        * Push these updated README files as an **additional commit** to your existing PR.

* **Updates from a Fork of the `cncf/toc` Repository:**
    * If you create a PR from a fork of the `cncf/toc` repository that modifies only the `tags.yaml` file, the `cncf-automation-bot` will:
        * Create a **new PR** on the `cncf/toc` repository.
        * This new PR will contain both your changes to the `tags.yaml` file and the automatically generated updates to the relevant README files.
        * The `cncf-automation-bot` will automatically close the original PR from your fork and add comments to both the original and the new PR to link them.

## **Key Takeaway:**

To update the README files in the `tags` and `toc_subprojects` directories, your PR **must only include modifications to the `tags.yaml` file**. The automation will then take care of generating and updating the corresponding README files.