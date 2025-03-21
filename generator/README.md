# TAG and TOC Subproject README Generator

This folder contains a Go program (`readme_app.go`) that automates the generation of README files for Technical Advisory Groups (TAGs) and TOC Subprojects. The README files are populated with data from the `tags.yaml` file.

## Purpose

The primary goals of this program are to:

* **Centralize Data:** Maintain all TAG and TOC Subproject information in a single, authoritative source (`tags.yaml`).
* **Automate README Creation:** Generate consistent and up-to-date README files, reducing manual effort and potential inconsistencies.
* **Streamline Updates:** Simplify the process of updating TAG and TOC Subproject information.

## How it Works

1.  **`tags.yaml` as the Data Source:**

    * The `tags.yaml` file in the project root holds all the essential details for each TAG and TOC Subproject. This includes:
        * Name
        * Mission statement
        * Leadership (Chairs, Tech Leads, Subproject Leads)
        * Meeting schedules
        * Contact information (Slack, mailing lists)
        * Subprojects
        * Charter link
    * **All updates** to TAG and TOC Subproject information should be made in this file.

2.  **README Generation with `readme_app.go`:**

    * The `readme_app.go` program reads the data from `tags.yaml`.
    * It uses Go templates (`tag_readme.tmpl` and `toc_subproject_readme.tmpl`) to format the data into README files.
    * For each TAG and TOC Subproject defined in `tags.yaml`, the program generates a corresponding `README.md` file in the appropriate directory.
    * It also generates a `charter.md` file if a charter link is provided in `tags.yaml`.

3.  **Directory Structure**

    * TAG README files are generated in the `tags/{tag.Dir}` directory.
    * TOC Subproject README files are generated in the `toc_subprojects/{toc_subproject.Dir}` directory.

## Workflow for Updating READMEs

The process for updating README files is automated using a GitHub workflow, with slight differences depending on the source of the Pull Request (PR):

1.  **Modify `tags.yaml`:** Make the necessary changes to the TAG or TOC Subproject entry in the `tags.yaml` file.
2.  **Create a Pull Request (PR):** Submit a PR with your changes to the `tags.yaml` file.
3.  **Automation Triggered:** The GitHub workflow is triggered by the PR. This workflow runs the `readme_app.go` program.
4.  **READMEs Updated:** The program generates the updated README files.
5.  **Handling PRs:**
    * **PR from a Local Branch:** If the PR comes from a branch within the same repository, the automation will add a new commit to the PR, containing the updated README files.
    * **PR from a Fork:** If the PR comes from a fork, the automation will create a *new* PR with the updated README files, and close the original PR.  The automation will also comment on both the original PR and the new PR to provide a link between them.
6.  **Review and Merge:**
    * **PR from a Local Branch:** Review the updated README files in the original PR. If everything is correct, merge the PR.
    * **PR from a Fork:** Review the *new* PR created by the automation. If the changes are correct, merge the *new* PR.

## `tags.yaml` Structure

```yaml
tags:
  - dir: "tag-example"
    name: "Example TAG"
    mission_statement: "This is an example TAG."
    leadership:
      chairs:
        - name: "Alex Johnson"
          github: "alexjohnson"
          company: "Example Corp" #optional
      tech_leads:
        - name: "Patrice Williams"
          github: "patricewilliams"
          company: "Another Corp"
      subproject_leads: # Added subproject leads
        - name: "Jordan Brown"
          github: "jordanbrown"
          company: "Some Company"
    meetings:
      - description: "TAG Meeting"
        tag_calendar: "[https://example.com/calendar](https://example.com/calendar)"
        recordings_url: "[https://example.com/recordings](https://example.com/recordings)"
    contact:
      slack: "#example-tag"
      mailing_list: "tag-example@lists.cncf.io"
      toc_liaison:
        name: "TOC Person"
        github: "tocperson"
        company: "TOC"
    tag_subprojects:
      - name: "Example Subproject"
        contact:
          slack: "#example-subproject"
          mailing_list: "subproject-example@lists.cncf.io"
    charter_link: "[https://github.com/cncf/tag-example/blob/main/charter.md](https://github.com/cncf/tag-example/blob/main/charter.md)"
toc_subprojects:
  - dir: "toc-subproject-example"
    name: "Example TOC Subproject"
    mission_statement: "This is an example TOC Subproject."
    leadership:
      chairs:
        - name: "Alex Johnson"
          github: "alexjohnson"
          company: "Example Corp"
      tech_leads:
        - name: "Patrice Williams"
          github: "patricewilliams"
          company: "Another Corp"
      subproject_leads:
        - name: "Jordan Brown"
          github: "jordanbrown"
          company: "Some Company"
    meetings:
      - description: "TOC Subproject Meeting"
        tag_calendar: "[https://example.com/calendar](https://example.com/calendar)"
        recordings_url: "[https://example.com/recordings](https://example.com/recordings)"
    contact:
      slack: "#toc-subproject-example"
      mailing_list: "toc-subproject-example@lists.cncf.io"
      toc_liaison:
        name: "TOC Person"
        github: "tocperson"
        company: "TOC"
    charter_link: "[https://github.com/cncf/toc-subproject-example/blob/main/charter.md](https://github.com/cncf/toc-subproject-example/blob/main/charter.md)"
