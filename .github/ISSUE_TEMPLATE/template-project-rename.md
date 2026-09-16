---
name: Project Rename Checklist
about: This template outlines the administrative steps required when a CNCF project is renamed.
title: "[RENAME] $OLD_PROJECT_NAME to $NEW_PROJECT_NAME"
labels:
- toc
- kind/rename
---

# Project Rename Checklist

This template outlines the administrative steps required when a CNCF project is renamed.

- **Current Project Name:** $OLD_PROJECT_NAME
- **New Project Name:** $NEW_PROJECT_NAME
- **New Project URL:** 
- **New GitHub Org:** 
- **Rename approval / TOC decision link:** 

---

## Rename Checklist

### GitHub
- [ ] Rename GitHub org (or create new GitHub org)
- [ ] Invite the new GH org to CNCF GHE account

### Website
- [ ] Acquire new web domain

### Formation team
- [ ] Request Formation team to:
  - [ ] Rename project in PCC
  - [ ] Update the project slug
  - [ ] Update the Github repo link
  - [ ] Update the website link

### LFX / PCC
- [ ] check PCC for updated name & PCC slug
- [ ] Connect (new) GitHub org to PCC
- [ ] Request PCC team to replace the mailing list slug
- [ ] Recreate the mailing list with the new name
  - [ ] Replace mailing list in the maintainers google mailing list
  

### Landscape
- [ ] Update project entry on the [CNCF Landscape](https://github.com/cncf/landscape/blob/master/landscape.yml)
  - [ ] Update landscape slug
  - [ ] Artwork url
  - [ ] web url
  - [ ] Any other relevant fields

### Monitoring & data
- [ ] Update project name in [CLOMonitor](https://github.com/cncf/clomonitor/blob/main/data/cncf.yaml)
  - [ ] Project name
  - [ ] Project slug
- [ ] Update [DevStats](https://github.com/cncf/devstats)
  - [ ] Project name

### .Project
- [ ] Rename project in .project automation workflow
- [ ] Rename project in CNCF Service Desk
- [ ] replace mailing list and mailing list number in .project maintainer automation workflow

### Artwork / branding
- [ ] Request a [creative brief](https://docs.google.com/document/d/1eL4HSMP2f4dkQCjl8-Kib05fpZb21ZjnkT2O12_jOl0/edit?tab=t.0#heading=h.w3zjcxsxtny6) from the project
- [ ] Open a [creative services ticket](https://form.asana.com/?k=n06lpzv4vdDJxUDR9L7SeQ&d=9283783873717) for the new logo
- [ ] Replace artwork in [cncf/artwork](https://github.com/cncf/artwork)
- [ ] Replace old artwork with new artwork
- [ ] Change the artwork hyperlink in the landscape entry

### Communications
- [ ] Rename Slack channels
- [ ] Notify CNCF PR / marketing team of the rename in slides and marketing materials

### Automation
- [ ] Run Calendar automation after PCC rename

### Fossa / Snyk
- [ ] Rename project

---

## Related links / references

- Original rename issue / discussion:
- Related PRs:
- Other:
